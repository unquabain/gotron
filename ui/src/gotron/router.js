import { useSelector, useDispatch } from 'react-redux'
import actions from '../actions'
import { navPathSelector } from '../selectors'

export function Route(props) {
  const navPath = useSelector(navPathSelector)
  const { path, children } = props
  // TODO: Make this fancier
  if (path !== navPath) {
    return (<></>)
  }
  return (<>{children}</>)
}

export function Link(props) {
  const dispatch = useDispatch()
  const navPath = useSelector(navPathSelector)
  const { path, children, className } = props
  const { session: { navigate } } = actions
  const classNames = [className]

  const handleClick = () => dispatch(navigate(path))
  classNames.push('gotron-link')
  if (path === navPath) {
    classNames.push('active')
  }

  return (
    <button type="button" className={classNames.join(' ')} href="#" onClick={handleClick}>{children}</button>
  )
}
