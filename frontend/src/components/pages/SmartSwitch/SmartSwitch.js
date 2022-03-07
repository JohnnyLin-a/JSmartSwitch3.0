import { useState } from "react";
import "./SmartSwitch.css"


const SmartSwitch = () => {
  const [confirmations, setConfirmations] = useState({});
  const [message, setMessage] = useState("");

  const resetConfirmation = key => {
    return () => {
      let currentState = { ...confirmations }
      delete currentState[key]
      setConfirmations(currentState)
    }
  }

  const misclickConfirm = (callback, key) => {
    return () => {
      const now = new Date().toLocaleTimeString("en-GB", { timeZone: "America/New_York" });
      const hour = parseInt(now.substr(0, now.indexOf(":")));
      
      
      if (hour <= 23 && hour >= 7) {
        callback()
        return
      }
      
      let currentState = { ...confirmations }
      if (typeof currentState[key] === 'undefined') {
        currentState[key] = {}
        setConfirmations(currentState)
        setTimeout(resetConfirmation(key), 3000)
      } else {
        resetConfirmation(key)()
        callback()
      }
    }
  }

  const openComputer = async () => {
    fetch('/api/v1/wol/wake', { method: 'POST'})
    setMessage("Opened computer")
    setTimeout(() => {
      setMessage("")
    }, 3000)
  }

  const openLights = async () => {
    fetch('/api/v1/lights/on', { method: 'POST'})
    setMessage("Opened lights")
    setTimeout(() => {
      setMessage("")
    }, 3000)
  }

  const closeLights = async () => {
    fetch('/api/v1/lights/off', { method: 'POST'})
    setMessage("Closed lights")
    setTimeout(() => {
      setMessage("")
    }, 3000)
  }

  return (
    <div style={{ display: "flex", flexDirection: "column", backgroundColor: "black", height: "100%" }}>
      <div className="super-center" style={{ height: "25%" }}>
        <h1 style={{ margin: 0, textAlign: "center", color: "white", fontSize: "8em" }}>Command Center</h1>
      </div>
      <div className="super-center flex-column" style={{ height: "55%" }}>
        <div className="m-3" style={{ width: "80%", height: "8rem" }}>
          <button className="btn btn-warning btn-block" style={{ width: "100%", height: "100%", fontSize: "3em" }} onClick={misclickConfirm(openComputer, "Open computer")}>{typeof confirmations["Open computer"] === 'undefined' ? "Open computer" : "Confirm?"}</button>
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
      <div style={{ textAlign: "center", color: "white", fontSize: "3em" }}>
        {message}
      </div>
    </div>
  )
}

export default SmartSwitch;
