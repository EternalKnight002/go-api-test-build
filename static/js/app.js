document.addEventListener("DOMContentLoaded", function () {
  const nameInput = document.getElementById("name");
  const btn = document.getElementById("sayHi");
  const result = document.getElementById("result");

  async function callGreet(name) {
    const url = name ? `/greet?name=${encodeURIComponent(name)}` : "/greet";
    try {
      result.textContent = "Loading…";
      const res = await fetch(url);
      if (!res.ok) {
        result.textContent = `Error: ${res.status} ${res.statusText}`;
        return;
      }
      const j = await res.json();
      result.textContent = j.message;
    } catch (err) {
      result.textContent = "Network or server error";
      console.error(err);
    }
  }

  btn.addEventListener("click", function () {
    callGreet(nameInput.value.trim());
  });

  // optional: auto-run once with default name
  // callGreet("");
});
