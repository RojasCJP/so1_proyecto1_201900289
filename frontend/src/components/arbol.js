import React from "react";
import { clientcpu } from "./cpu";
export class Arbol extends React.Component {
    client = clientcpu
    state = {
        cpuData: { Processes: [], Running: 0, Sleeping: 0, Zombie: 0, Stopped: 0, Total: 0, Usage: 0 }
    }

    componentDidMount() {
        this.client.onopen = (event) => {
            console.log("cpu websocket connected")
        }
        this.client.onmessage = (message) => {
            const dataFromServer = JSON.parse(message.data)
            console.log("cpu arbol", dataFromServer)
            this.setState({ cpuData: dataFromServer })
        }
    }

    render() { 
        return(
            <div className='col'>
                    <p>
                        Running: {this.state.cpuData.Running}
                    </p>
                    <p>
                        Sleeping: {this.state.cpuData.Sleeping}
                    </p>
                    <p>
                        Zombie: {this.state.cpuData.Zombie}
                    </p>
                    <p>
                        Stopped: {this.state.cpuData.Stopped}
                    </p>
                    <p>
                        Total: {this.state.cpuData.Total}
                    </p>
                    <p>
                        Usage: {this.state.cpuData.Usage}
                    </p>
                </div>
        )
    }
}