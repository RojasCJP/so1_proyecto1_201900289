import React from 'react'
import { Chart } from 'react-google-charts'


const client = new WebSocket('ws://localhost:4200/cpu')
export class Cpu extends React.Component {
    state = {
        data: [['x', 'Memoria RAM'], [1, 0], [2, 0], [3, 0], [4, 0], [5, 0], [6, 0], [7, 0], [8, 0], [9, 0], [10, 0], [11, 0], [12, 0], [13, 0], [14, 0], [15, 0]],
        cpuData:{Processes:[],Running:0,Sleeping:0,Zombie:0,Stopped:0,Total:0,Usage:0}
    }

    componentDidMount() {
        client.onopen = (event) => {
            console.log("memory websocket connected");
        }
        client.onmessage = (message) => {
            const dataFromServer = JSON.parse(message.data);
            console.log("cpu", dataFromServer)
            this.fillData()
            this.setState({cpuData:dataFromServer})
        }
    }

    componentWillUnmount() {
        client.close()
    }

    render() {
        return (
            <div className='row'>
            <div className='col'>

                <Chart
                    width={'800px'}
                    height={'1000px'}
                    chartType="LineChart"
                    loader={<div>Loading Chart</div>}
                    data={this.state.data}
                    options={{
                        title: ' ',
                        backgroundColor: 'transparent',
                        hAxis: {
                            title: 'Tiempo',
                            textPosition:'none'
                        },
                        vAxis: {
                            title: 'Uso',
                            minValue: 0,
                            maxValue: 100
                        },
                    }}

                    rootProps={{ 'data-testid': '1' }}
                />
            </div>
                    <div className='col'>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
                        <br/>
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
        </div>
        )
    }

    fillData() {
        var encabezado = ['x', 'Memoria RAM']
        var inputData = [Number(this.state.data[15][0]) + 1, this.state.cpuData.Usage]
        console.log(this.state.data[7])
        var datos = []
        datos.push(encabezado)
        for (let i = 0; i < 15; i++) {
            if(this.state.data[i+2]){
                datos.push(this.state.data[i+2])
            }
        }
        datos.push(inputData)
        this.setState({ data: datos })
    }
}