import logo from './logo.svg'
import './App.css'
import { useSelector, useDispatch } from 'react-redux'
import actions from './actions'
import { usernameSelector } from './selectors'
import { Route, Link } from './gotron_router'

function App() {
  const name = useSelector(usernameSelector)
  const dispatch = useDispatch()
  return (
    <div className="App">
      <nav>
        <Link path="/">Home</Link>
        <Link path="/fortune">Fortune</Link>
      </nav>
          <Route path="/">
          <h1>Hello from GoTron</h1>
          <p>My name is {name}.</p>
          <input
            type="text"
            defaultValue={name}
            onChange={({target: { value } }) => dispatch(actions.session.updateUsername(value))}
          />
          </Route>
          <Route path="/fortune">
          <pre>{`
write-protect tab, n.:
	A small sticker created to cover the unsightly notch carelessly left
	by disk manufacturers.  The use of the tab creates an error message
	once in a while, but its aesthetic value far outweighs the momentary
	inconvenience.
		-- Robb Russon
          `}</pre>
          </Route>

    </div>
  );
}

export default App;
