import "./SmartSwitch.css"


const SmartSwitch = () => {

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
    <div style={{ display: "flex", flexDirection: "column", backgroundColor: "black", height: "100%" }}>
      <div className="super-center" style={{ height: "25%" }}>
        <h1 style={{ margin: 0, textAlign: "center", color: "white", fontSize: "8em" }}>Command Center</h1>
      </div>
      <div className="super-center flex-column" style={{ height: "75%" }}>
        <div className="m-3" style={{ width: "80%", height: "8rem" }}>
          <button className="btn btn-warning btn-block" style={{ width: "100%", height: "100%", fontSize: "3em" }} onClick={openComputer}>Open computer</button>
        </div>
        <div className="m-3 d-flex flex-row" style={{ width: "80%", height: "8rem" }}>
          <div style={{ paddingRight: "1rem", width: "50%" }}>
            <button className="btn btn-success btn-block super-max" style={{ fontSize: "3em" }} onClick={openLights}>Open lights</button>
          </div>
          <div style={{ paddingLeft: "1rem", width: "50%" }}>
            <button className="btn btn-danger btn-block super-max" style={{ fontSize: "3em" }} onClick={closeLights}>Close lights</button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default SmartSwitch;
