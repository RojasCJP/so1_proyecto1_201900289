import React from "react";


const client = new WebSocket('ws://localhost:4200/cpu')

export class Process extends React.Component {
    state = {
        cpuData: { Processes: [], Running: 0, Sleeping: 0, Zombie: 0, Stopped: 0, Total: 0, Usage: 0 }
    }

    componentDidMount() {
        client.onopen = (event) => {
            console.log("memory websocket connected");
        }
        client.onmessage = (message) => {
            const dataFromServer = JSON.parse(message.data);
            console.log("cpu", dataFromServer)
            this.fillData()
            this.setState({ cpuData: dataFromServer })
        }
    }

    componentWillUnmount() {
        client.close()
    }

    render() {
        return (<div>
            <h1>Procesos</h1>
            <br/>
            <div className='row'>
                <div className="col">Running: {this.state.cpuData.Running}</div>
                <div className="col">Sleeping: {this.state.cpuData.Sleeping}</div>
                <div className="col">Zombie: {this.state.cpuData.Zombie}</div>
                <div className="col">Stopped: {this.state.cpuData.Stopped}</div>
                <div className="col">Total: {this.state.cpuData.Total}</div>
                <div className="col">Usage: {this.state.cpuData.Usage}</div>
            </div>
            <br/>
            <table className="table table-striped table-hover table-light">
                <thead>
                    <tr>
                        <th scope='col'>PID</th>
                        <th scope='col'>Name</th>
                        <th scope='col'>User</th>
                        <th scope='col'>State</th>
                        <th scope='col'>RAM</th>
                    </tr>
                </thead>
                <tbody style={{textAlign:'center'}}>
                    {this.state.cpuData.Processes.map(
                        element =>
                        <tr key={element.Pid}>
                            <td>{element.Pid}</td>
                            <td>{element.Name}</td>
                            <td>{element.User}</td>
                            <td>{element.State}</td>
                            <td>{element.Ram}</td>
                        </tr>
                    )}
                </tbody>
            </table>
        </div>)
    }
}