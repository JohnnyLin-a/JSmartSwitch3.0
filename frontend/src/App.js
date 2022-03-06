
function App() {

  const openComputer = async () => {
    fetch('/api/v1/wol/wake', { method: 'POST'})
  }

  const openLights = async () => {
    fetch('/api/v1/lights/on', { method: 'POST'})
  }

  const closeLights = async () => {
    fetch('/api/v1/lights/off', { method: 'POST'})
  }

  return (
    <div style={{ display: "flex", flexDirection: "column", backgroundColor: "black", minHeight: "100%" }}>
      <header style={{ display: "flex", flex: 1, alignItems:"center" }}>
        <div className="container">
          <h1 style={{ textAlign: "center", color: "white", fontSize: "8em" }}>Command Center</h1>
        </div>
      </header>
      <div style={{ display: "flex", flex: 3, flexDirection: "column", alignItems: "center"}}>
        <div style={{ justifyContent: "center"}}>
          <div className="container" style={{ }}>
            <button className="btn btn-warning" style={{ }} onClick={openComputer}>Power on computer</button>
          </div>
          <div className="container" style={{ }}>
            <button className="btn btn-success" onClick={openLights}>Open lights</button>
            <button className="btn btn-danger"onClick={closeLights}>Close lights</button>
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;
