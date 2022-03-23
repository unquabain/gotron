import logo from './logo.svg'
import './App.css'
import { useSelector, useDispatch } from 'react-redux'
import actions from './actions'
import { usernameSelector, fortuneSelector } from './selectors'
import { Route, Link } from './gotronRouter'

function App() {
  const name = useSelector(usernameSelector)
  const fortune = useSelector(fortuneSelector)
  const dispatch = useDispatch()
  const { session: { updateUsername }, views: { fortune: { generate: generateFortune } } } = actions
  return (
    <div className="App">
      <nav>
        <Link path="/">Home</Link>
        <Link path="/fortune">Fortune</Link>
        <Link path="/session" classNames={['right']}>{name}</Link>
      </nav>
          <Route path="/">
            <h1>Welcome to the GoTron example application</h1>
            <p>GoTron is an experimental project to use React technologies to provide an
               interface to a Go application without opening ports on the target computer.
               All interaction between the Go application and the front-end is done via
               a custom Redux "store" object that is an interface to bound Go functions.
            </p>

            <p>
               The <code>dispatch()</code> function sends the method directly to the Go
               application code, and when Go changes its internal state, it pushes that
               to the front-end application stack to trigger updated displays.
            </p>

            <p>
               By default, the state is organized into three sections:
            </p>
            <ol>
              <li><code>session</code> contains information about the current user and the
                  running application, such as the `navPath`, which chooses the current view.</li>
              <li><code>data</code> contains the canonical application data.</li>
              <li><code>views</code> contains data about each view. Not the purely presentational
                  data, which should be maintained by the component, but for example, a record that
                  the view is building that it needs to keep track of before it commits; or a pre-filtered
                  subset of some data pulled from the <code>data</code> section.</li>
            </ol>
            <p>
               Each of these sections should be statically-typed on the Go side for consistency and
               compile-time provability.
            </p>
          </Route>
          <Route path="/session">
          <h1>Hello from GoTron</h1>
          <p>My name is {name}.</p>
          <input
            type="text"
            defaultValue={name}
            onChange={({target: { value } }) => dispatch(updateUsername(value))}
          />
          </Route>
          <Route path="/fortune">
          <pre>{fortune}</pre>
          <button type="button" className={'btn'} onClick={()=>dispatch(generateFortune())}>Generate</button>
          </Route>

    </div>
  );
}

export default App;
