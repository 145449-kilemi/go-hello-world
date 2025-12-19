// -------------------------
// Dark/Light Mode Toggle
// -------------------------
const themeToggle = document.getElementById("themeToggle");
themeToggle.addEventListener("click", () => {
  document.body.classList.toggle("dark-mode");
});

// -------------------------
// Collapsible Sections
// -------------------------
document.querySelectorAll(".collapsible h2").forEach(header => {
  header.addEventListener("click", () => {
    const card = header.parentElement;
    card.classList.toggle("collapsed");
    const toggleBtn = header.querySelector(".toggleBtn");
    toggleBtn.textContent = card.classList.contains("collapsed") ? "►" : "▼";
  });
});

// -------------------------
// Dynamic List
// -------------------------
async function loadList() {
  const res = await fetch("/api/data");
  const data = await res.json();
  const itemList = document.getElementById("itemList");
  itemList.innerHTML = "";
  data.items.forEach(item => {
    const li = document.createElement("li");
    li.textContent = item;
    itemList.appendChild(li);
  });
}
loadList();

// -------------------------
// Form Submission
// -------------------------
const form = document.getElementById("myForm");
const formResponse = document.getElementById("formResponse");
form.addEventListener("submit", async (e) => {
  e.preventDefault();
  const formData = new FormData(form);
  const res = await fetch("/api/submit", {
    method: "POST",
    body: formData
  });
  const data = await res.json();
  formResponse.textContent = data.message;
  form.reset();
});

// -------------------------
// Live Messages (Chat)
// -------------------------
const liveUpdates = document.getElementById("liveUpdates");
const messageForm = document.getElementById("messageForm");
const messageInput = document.getElementById("messageInput");
const maxMessages = 20;

// SSE connection for live messages
const evtSource = new EventSource("/stream");
evtSource.onmessage = function(event) {
  const li = document.createElement("li");
  li.textContent = event.data;
  liveUpdates.appendChild(li);

  while (liveUpdates.children.length > maxMessages) {
    liveUpdates.removeChild(liveUpdates.firstChild);
  }

  liveUpdates.scrollTop = liveUpdates.scrollHeight;
};

// Submit new message
messageForm.addEventListener("submit", async (e) => {
  e.preventDefault();
  const msg = messageInput.value.trim();
  if (!msg) return;

  await fetch("/send", {
    method: "POST",
    headers: { "Content-Type": "application/x-www-form-urlencoded" },
    body: "message=" + encodeURIComponent(msg)
  });

  messageInput.value = "";
});

// -------------------------
// Live Line Chart
// -------------------------
let lineChart;
let evtSourceLine;

function startLineChart() {
  const lineCtx = document.getElementById("liveLineChart").getContext("2d");
  if (lineChart) lineChart.destroy();

  lineChart = new Chart(lineCtx, {
    type: 'line',
    data: {
      labels: [],
      datasets: [{
        label: 'Live Line',
        data: [],
        borderColor: 'blue',
        fill: false
      }]
    },
    options: { animation: false }
  });

  if (evtSourceLine) evtSourceLine.close();
  evtSourceLine = new EventSource("/chart-stream-line");
  evtSourceLine.onmessage = function(event) {
    const data = JSON.parse(event.data);
    lineChart.data.labels.push(data.label);
    lineChart.data.datasets[0].data.push(data.value);

    if (lineChart.data.labels.length > 20) {
      lineChart.data.labels.shift();
      lineChart.data.datasets[0].data.shift();
    }

    lineChart.update();
  };
}

document.getElementById("startLine").addEventListener("click", startLineChart);
document.getElementById("stopLine").addEventListener("click", () => {
  if (evtSourceLine) evtSourceLine.close();
});

// -------------------------
// Live Bar Chart
// -------------------------
let barChart;
let evtSourceBar;

function startBarChart() {
  const barCtx = document.getElementById("liveBarChart").getContext("2d");
  if (barChart) barChart.destroy();

  barChart = new Chart(barCtx, {
    type: 'bar',
    data: {
      labels: [],
      datasets: [{
        label: 'Live Bar',
        data: [],
        backgroundColor: '#ff7f50'
      }]
    },
    options: { animation: false }
  });

  if (evtSourceBar) evtSourceBar.close();
  evtSourceBar = new EventSource("/chart-stream-bar");
  evtSourceBar.onmessage = function(event) {
    const data = JSON.parse(event.data);
    barChart.data.labels.push(data.label);
    barChart.data.datasets[0].data.push(data.value);

    if (barChart.data.labels.length > 20) {
      barChart.data.labels.shift();
      barChart.data.datasets[0].data.shift();
    }

    barChart.update();
  };
}

document.getElementById("startBar").addEventListener("click", startBarChart);
document.getElementById("stopBar").addEventListener("click", () => {
  if (evtSourceBar) evtSourceBar.close();
});
