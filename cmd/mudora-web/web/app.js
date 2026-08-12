const status = document.getElementById("status");
const permalink = document.getElementById("permalink");
const startScreenCode = document.getElementById("startscreen-code");
const romInput = document.getElementById("rom-input");
const search = document.getElementById("search");
const searchBtn = document.getElementById("search-button");
const clearBtn = document.getElementById("clear");
const results = document.getElementById("results");
const playthrough = document.getElementById("playthrough");
const version = document.getElementById("version");

let romBytes = null;
let ready = false;

const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject)
  .then((result) => {
    go.run(result.instance);
    ready = true;
    version.textContent = "v" + window.mudoraVersion;
    status.textContent = "Choose a ROM file. It is parsed entirely in your browser and never uploaded.";
  })
  .catch((err) => {
    status.textContent = "Failed to load WASM module: " + err;
  });

romInput.addEventListener("change", async () => {
  const file = romInput.files[0];
  if (!file) return;
  romBytes = new Uint8Array(await file.arrayBuffer());
  search.disabled = false;
  status.textContent = `Inspecting ${file.name}`;
  render(search.value);
});

search.addEventListener("input", () => {
  console.log(search.value);
  const disabled = search.value.trim() === "";
  searchBtn.disabled = disabled;
  clearBtn.disabled = disabled;
});
searchBtn.addEventListener("click", () => {
  render(search.value);
})
clearBtn.addEventListener("click", () => {
  search.value = "";
  render("");
});

function render(query) {
  if (!ready || !romBytes) return;

  const permalinkRaw = window.mudoraPermalink(romBytes);
  const parsedPermalink = tryParseJSON(permalinkRaw);
  if (parsedPermalink && parsedPermalink.error) {
    permalink.textContent = "Unable to resolve permalink.";
  } else {
    permalink.textContent = parsedPermalink.permalink;
    permalink.href = parsedPermalink.permalink;
  }
  document.getElementById("permalink-container").style.display = "";

  const startscreenCodeRaw = window.mudoraItemHash(romBytes);
  const parsedCode = tryParseJSON(startscreenCodeRaw);
  if (parsedCode && !parsedCode.error) {
    startScreenCode.textContent = "";
    for (const item of parsedCode) {
      startScreenCode.appendChild(makeIcon(item.icon, item.name));
    }
  } else {
    startScreenCode.textContent = "Unable to parse ROM code.";
  }
  startScreenCode.style.display = "";

  const raw = window.mudoraInspect(romBytes, query);
  const parsed = JSON.parse(raw);
  if (parsed && parsed.error) {
    status.textContent = "Error: " + parsed.error;
    return;
  }

  function createHeader(name, expandOrCollapseAllFn) {
    const headingContainer = document.createElement("div");
    headingContainer.className = "heading-container";

    const heading = document.createElement("h1");
    heading.innerText = name;
    headingContainer.appendChild(heading);

    const expandOrCollapseAllButton = document.createElement("button");
    expandOrCollapseAllButton.innerText = "Expand/Collapse All";
    expandOrCollapseAllButton.onclick = (e) => {
      e.preventDefault();
      expandOrCollapseAllFn();
    };
    headingContainer.appendChild(expandOrCollapseAllButton);

    return headingContainer;
  }

  function expandOrCollapseAll(parent) {
    const regions = parent.querySelectorAll(".region");

    let allExpanded = true;

    for (const region of regions) {
      if (!region.classList.contains("expanded")) {
        allExpanded = false;
        break;
      }
    }

    for (const region of regions) {
      if (allExpanded) {
        region.classList.remove("expanded");
      } else {
        region.classList.add("expanded");
      }
    }
  }

  results.innerHTML = "";
  results.appendChild(createHeader("All Locations", () => expandOrCollapseAll(results)));

  const collapsed = query.trim() === "";
  for (const group of parsed) {
    results.appendChild(buildRegion(group, collapsed));
  }

  const rawPlaythrough = window.mudoraSolve(romBytes);
  const parsedPlaythrough = JSON.parse(rawPlaythrough);

  if (parsedPlaythrough && parsedPlaythrough.error) {
    status.textContent = "Error: " + parsedPlaythrough.error;
    return;
  }

  const hr = document.getElementById("results-playthrough-divider");
  hr.removeAttribute("hidden");

  playthrough.innerHTML = "";
  playthrough.appendChild(createHeader("Playthrough - Shortest Path", () => expandOrCollapseAll(playthrough)));

  let step = 0;
  while (parsedPlaythrough[step] !== undefined) {   
    const locations = [];
    for (const locId in parsedPlaythrough[step]) {
      locations.push(parsedPlaythrough[step][locId]);
    }

    playthrough.appendChild(buildStep(step + 1, locations, collapsed));

    step += 1;
  }
}

function buildRegion(group, collapsed) {
  const section = document.createElement("div");
  section.className = "region" + (collapsed ? "" : " expanded");

  const header = document.createElement("div");
  header.className = "region-header";
  header.addEventListener("click", () => section.classList.toggle("expanded"));

  const dungeonReward = document.createElement("div");
  dungeonReward.className = "region-dungeon-reward";
  header.appendChild(dungeonReward);

  const name = document.createElement("span");
  name.className = "region-name";
  name.textContent = group.region;
  header.appendChild(name);

  for (const loc of group.locations) {
    if (loc.location.includes("Prize")) {
      dungeonReward.appendChild(makeIcon(loc.icon, loc.item));
    } else if (loc.progression && loc.icon) {
      header.appendChild(makeIcon(loc.icon, loc.item));
    }
  }

  const rows = document.createElement("div");
  rows.className = "region-rows";
  for (const loc of group.locations) {
    rows.appendChild(buildRow(loc));
  }

  section.appendChild(header);
  section.appendChild(rows);
  return section;
}

function buildStep(step, locations, collapsed) {
  const section = document.createElement("div");
  section.className = "region" + (collapsed ? "" : " expanded");

  const header = document.createElement("div");
  header.className = "region-header";
  header.addEventListener("click", () => section.classList.toggle("expanded"));

  const name = document.createElement("span");
  name.className = "region-name";
  name.textContent = `Step ${step}`;
  header.appendChild(name);

  const rows = document.createElement("div");
  rows.className = "region-rows";
  for (const loc of locations) {
    rows.appendChild(buildRow(loc));
  }

  section.appendChild(header);
  section.appendChild(rows);
  return section;
}

function buildRow(loc) {
  const row = document.createElement("div");
  row.className = "row";

  const locLabel = document.createElement("span");
  locLabel.textContent = loc.location;

  const icon = loc.icon ? makeIcon(loc.icon, loc.item) : document.createElement("span");

  const itemLabel = document.createElement("span");
  itemLabel.textContent = loc.item;

  row.appendChild(locLabel);
  row.appendChild(icon);
  row.appendChild(itemLabel);
  return row;
}

function makeIcon(src, alt) {
  const img = document.createElement("img");
  img.className = "icon";
  img.src = src;
  img.alt = alt;
  return img;
}

function tryParseJSON(raw) {
  try {
    parsed = JSON.parse(raw)

    return parsed
  } catch (e) {
    console.error("Error parsing JSON", e)
    return { error: e }
  }
}
