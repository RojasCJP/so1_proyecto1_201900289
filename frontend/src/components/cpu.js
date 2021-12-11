import React from 'react'
import { Chart } from 'react-google-charts'


const client = new WebSocket('ws://localhost:4200/cpu')
export class Cpu extends React.Component {
    state = {
        data: [['x', 'Memoria RAM'], [1, 2], [2, 3], [3, 5], [4, 7], [5, 8], [6, 11], [7, 1]]
    }

    componentDidMount() {
        client.onopen = (event) => {
            console.log("memory websocket connected");
        }
        client.onmessage = (message) => {
            const dataFromServer = JSON.parse(message.data);
            console.log("cpu", dataFromServer)
        }
    }

    componentWillUnmount() {
        client.close()
    }

    render() {
        return (
            <div>
                <Chart
                    width={'800px'}
                    height={'500px'}
                    chartType="LineChart"
                    loader={<div>Loading Chart</div>}
                    data={this.state.data}
                    options={{
                        title: ' ',
                        backgroundColor: 'transparent',
                        hAxis: {
                            title: 'Tiempo',
                        },
                        vAxis: {
                            title: 'Uso',
                        },
                    }}

                    rootProps={{ 'data-testid': '1' }}
                />
            </div>
        )
    }
}