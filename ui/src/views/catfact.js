import { useSelector, useDispatch } from 'react-redux'
import actions from '../actions'
import { catfactSelector } from '../selectors'
import View from './view'

function CatFact() {
  const catfact = useSelector(catfactSelector)
  const dispatch = useDispatch()
  const { views: { catfact: { fetch } } } = actions
  const catFactHandler = () => dispatch(fetch())
  return (<View>
    <p>This view demonstrates fetching data from an API asynchronously.</p>
    <p>Press the button below, and a fact about cats will be downloaded from the Internet.</p>
    <button type="button" onClick={catFactHandler}>Get a Cat Fact</button>
    <p>{catfact}</p>
  </View>)
}

export default CatFact
