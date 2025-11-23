document.addEventListener("DOMContentLoaded", function () {
  // Greet API elements
  const nameInput = document.getElementById("name");
  const btn = document.getElementById("sayHi");
  const result = document.getElementById("result");

  //  Double API elements
  const numberInput = document.getElementById("numberInput");
  const doubleBtn = document.getElementById("doDouble");
  const doubleResult = document.getElementById("doubleResult");

  //  Subtraction API elements
  const n1Input = document.getElementById("n1Input");
  const n2Input = document.getElementById("n2Input");
  const subtractBtn = document.getElementById("doSubtract");
  const subtractResult = document.getElementById("subtractResult");

  //  Pet Name API elements
  const adjectiveInput = document.getElementById("adjectiveInput");
  const animalInput = document.getElementById("animalInput");
  const petNameBtn = document.getElementById("getPetName");
  const petNameResult = document.getElementById("petNameResult");


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

  //  Function to call the /subtract API
  async function callSubtract(n1, n2) {
    const url = `/subtract?n1=${encodeURIComponent(n1)}&n2=${encodeURIComponent(n2)}`;
    try {
      subtractResult.textContent = "Loading…";
      const res = await fetch(url);
      if (!res.ok) {
        subtractResult.textContent = `Error: ${res.status} ${res.statusText}`;
        return;
      }
      const j = await res.json();
      // Display the result from the API
      subtractResult.textContent = `${j.n1} minus ${j.n2} is ${j.result}.`;
    } catch (err) {
      subtractResult.textContent = "Network or server error";
      console.error(err);
    }
  }

  // Function to call the /petname API (FIXED: JSON keys are now correctly accessed as j.animal and j.petName)
  async function callPetName(adjective, animal) {
    const adj = adjective.trim();
    const ani = animal.trim();
    // Use default values if inputs are empty for the URL
    const url = `/petname?adjective=${encodeURIComponent(adj || 'mysterious')}&animal=${encodeURIComponent(ani || 'shrimp')}`;
    try {
      petNameResult.textContent = "Loading…";
      const res = await fetch(url);
      if (!res.ok) {
        petNameResult.textContent = `Error: ${res.status} ${res.statusText}`;
        return;
      }
      const j = await res.json();
      // Display the result from the API
      // FIX: Changed j.Animal to j.animal and j.PetName to j.petName
      petNameResult.textContent = `Your new pet name for the ${j.animal} is: ${j.petName}`;
    } catch (err) {
      petNameResult.textContent = "Network or server error";
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

  // Event listener for the Subtraction API button
  subtractBtn.addEventListener("click", function () {
    callSubtract(n1Input.value.trim(), n2Input.value.trim());
  });

  // Event listener for the Pet Name API button
  petNameBtn.addEventListener("click", function () {
    callPetName(adjectiveInput.value.trim(), animalInput.value.trim());
  });

  // optional: auto-run once with default name
  // callGreet("");
});