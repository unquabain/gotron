window.gotronDispatch = window.gotronDispatch || ((action) => console.log('original gotronDispatch called', action))
window.gotronLog = window.gotronLog || console.log
window.gotronPushState = window.gotronPushState || (() => console.log('original pushState called'))

class GotronStore {
  constructor() {
    this.listeners = {}
    this.listenerID = 0
    this.state = {}
    const self = this
    window.gotronSetState = (newState) => {
      self.state = newState
      Object.values(this.listeners).forEach(val => val())
    }
    window.gotronPushState()
  }

  getState = () => this.state 

  subscribe = (listener) => {
    const id = this.listenerID
    this.listenerID++
    this.listeners[id] = listener
    const self = this 
    return () => delete self.listeners[id]
  }

  dispatch = (action) => window.gotronDispatch(action)

  replaceReducer = (nextReducer) => {
    throw('replaceReducer is not implemented')
  }
}

export default new GotronStore()
