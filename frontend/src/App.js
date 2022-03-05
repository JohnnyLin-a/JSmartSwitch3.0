
function App() {

  const openComputer = async () => {
    console.log("openComputer")
    fetch('/api/v1/wol/wake', { method: 'POST'})
      .then(response => response.json())
      .then(data => console.log(data));
  }

  const openLights = async () => {
    console.log("openLights")
    fetch('/api/v1/lights/on', { method: 'POST'})
      .then(response => response.json())
      .then(data => console.log(data));
  }

  const closeLights = async () => {
    console.log("closeLights")
    fetch('/api/v1/lights/off', { method: 'POST'})
      .then(response => response.json())
      .then(data => console.log(data));
  }

  return (
    <div style={{ backgroundColor: "black"}}>
      <header>
        <p>Year 2000 edition</p>
      </header>
      <button onClick={openComputer}>Power on computer</button><br/><br/>
      <button onClick={openLights}>Open lights</button><br/><br/>
      <button onClick={closeLights}>Close lights</button>
    </div>
  );
}

export default App;
