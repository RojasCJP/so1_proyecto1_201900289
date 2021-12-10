import React from 'react'
import {Chart} from 'react-google-charts'
import {w3cwebsocket as W3CWebSocket} from 'websocket'

const client = new W3CWebSocket('ws://localhost:4200/memory')

export class Memory extends React.Component {
    state = {
        data: [['x','Memoria RAM'],[1,2],[2,3],[3,5],[4,7],[5,8],[6,11],[7,1]]
    }
    
    componentDidMount(){
        client.onopen()=()=>{
            console.log("memory websocket connected");
        }
        client.onmessage = (message)=>{
            const dataFromServer = JSON.parse(message.data);
            console.log("data", dataFromServer)
        }
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
                        backgroundColor:'transparent',
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

    fillData() {
        var Data = {
            labels: ["Lunes", "Martes", "Miercoles", "Jueves", "Viernes"],
        }
        this.setState({ data: Data })
    }
}