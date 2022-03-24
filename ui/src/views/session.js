import { useSelector, useDispatch } from 'react-redux'
import { usernameSelector } from '../selectors'
import actions from '../actions'
import View from './view'

function Session() {
  const name = useSelector(usernameSelector)
  const dispatch = useDispatch()
  const { session: { updateUsername } } = actions
  return (<View>
    <h1>{name}</h1>
    <p>
    This view demonstrates a very simple form. It is remarkable in how unremarkable it is.
    </p>
    <p>
    The change events that are generated when you type in the box are
    dispatched to the Go program, where the statically-typed state
    is updated, and the listeners in JavaScript are notified. The
    round-trip is very rapid.
    </p>
    <input
      type="text"
      defaultValue={name}
      onChange={({target: { value } }) => dispatch(updateUsername(value))}
    />
  </View>)
}
export default Session
