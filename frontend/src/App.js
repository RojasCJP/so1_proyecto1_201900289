import logo from './logo.svg';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Welcome } from './components/welcome';
import { Cpu } from './components/cpu';
import { Memory } from './components/memory';
import { Navbar } from './components/navbar';
import { Process } from './components/process';
import history from './history/history';
import './App.css';


function App() {
  return (
    <div className="App">
      <Router history={history}>
        <Navbar></Navbar>
        <Routes>
          <Route exact path="/" element={<Welcome />} />
          <Route exact path="/cpu" element={<Cpu />} />
          <Route exact path="/memory" element={<Memory />} />
          <Route exact path="/process" element={<Process />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
