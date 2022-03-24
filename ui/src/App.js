import { useSelector } from 'react-redux'
import { usernameSelector } from './selectors'
import './App.css'
import { Route, Link } from './gotron/router'
import Fortune from './views/fortune'
import Home from './views/home'
import Session from './views/session'
import CatFact from './views/catfact'

function App() {
  const name = useSelector(usernameSelector)
  return (
    <div className="App">
      <nav>
        <Link path="/">Home</Link>
        <Link path="/fortune">Fortune</Link>
        <Link path="/catfact">Cat Fact</Link>
        <Link path="/session" className="right">{name}</Link>
      </nav>
      <Route path="/"><Home/></Route>
      <Route path="/fortune"><Fortune /></Route>
      <Route path="/catfact"><CatFact /></Route>
      <Route path="/session"><Session/></Route>
    </div>
  );
}

export default App;
