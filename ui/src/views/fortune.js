import { useSelector, useDispatch } from 'react-redux'
import actions from '../actions'
import { fortuneSelector } from '../selectors'
import View from './view'

function Fortune() {
  const fortune = useSelector(fortuneSelector)
  const { views: { fortune: { generate: generateFortune } } } = actions
  const dispatch = useDispatch()
  return (<View>
    <p>
    This view demonstrates using Go-side logic. The button below causes the compiled code to
    shell-out to the <code>fortune</code> program on your computer, if it exists, and populate
    the text area below.
    </p>
    <p>
    If <code>fortune</code> is not installed, the error message will appear below.
    </p>
    <button type="button" className={'btn'} onClick={()=>dispatch(generateFortune())}>Generate</button>
    <pre>{fortune}</pre>
  </View>)
}
export default Fortune
