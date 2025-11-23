document.addEventListener("DOMContentLoaded", function () {
  // Greet API elements
  const nameInput = document.getElementById("name");
  const btn = document.getElementById("sayHi");
  const result = document.getElementById("result");

  //  Double API elements
  const numberInput = document.getElementById("numberInput");
  const doubleBtn = document.getElementById("doDouble");
  const doubleResult = document.getElementById("doubleResult");


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

  //  Function to call the /double API
  async function callDouble(number) {
    const url = `/double?number=${encodeURIComponent(number)}`;
    try {
      doubleResult.textContent = "Loading…";
      const res = await fetch(url);
      if (!res.ok) {
        doubleResult.textContent = `Error: ${res.status} ${res.statusText}`;
        return;
      }
      const j = await res.json();
      // Display the result from the API
      doubleResult.textContent = `The double of ${j.input} is ${j.result}.`;
    } catch (err) {
      doubleResult.textContent = "Network or server error";
      console.error(err);
    }
  }


  btn.addEventListener("click", function () {
    callGreet(nameInput.value.trim());
  });

  //  Event listener for the Double API button
  doubleBtn.addEventListener("click", function () {
    callDouble(numberInput.value.trim());
  });

  // optional: auto-run once with default name
  // callGreet("");
});