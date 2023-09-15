<script>
  export let interfaceObj;
  export let characteristics;

  let formData = {};

  function renderInputs(obj, parentKey = '') {
    let renderedInputs = [];

    for (const key in obj) {
      const currentKey = parentKey ? `${parentKey}.${key}` : key;

      if (typeof obj[key] === 'object') {
        renderedInputs = [
          ...renderedInputs,
          ...renderInputs(obj[key], currentKey),
        ];
      } else {
        renderedInputs.push({
          key: currentKey,
          type: typeof obj[key],
        });
      }
    }

    return renderedInputs;
  }

  function convertToJson() {
    formData = {};
    const inputs = document.querySelectorAll('.include-in-json');

    inputs.forEach((input) => {
      formData[input.name] = input.value;
    });

    const jsonString = JSON.stringify(formData, null, 2);
    characteristics = jsonString;
  }
</script>

<div>
  {#each renderInputs(interfaceObj) as { key, type }}
  <label for={key}>{key}:</label>
  <input
    class="include-in-json bg-surface-900 border-neutral-900 shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline"
    type={type}
    id={key}
    name={key}
    required
  /><br />
  {/each}
  <button class="btn" on:click={convertToJson}>Convert to JSON</button>
</div>
