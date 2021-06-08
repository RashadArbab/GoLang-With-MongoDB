import logo from './logo.svg';
import './App.css';
import axios from 'axios'; 
import {useEffect, useState} from 'react'; 

function App() {

  const [output, setOutput] = useState("")

  useEffect(()=>{
    axios.get("/api/")
    .then((res)=>{
      setOutput(res.data)
    })
  }, [])

  return (
    <div className="App">
      this is the output: 
      {JSON.stringify(output)}
    </div>
  );
}

export default App;
