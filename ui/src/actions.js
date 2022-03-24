const actions = {
  session: {
    updateUsername: un => ({ action: 'SESSION.UPDATE_USERNAME', payload: un }),
    navigate: path => ({ action: 'SESSION.NAVIGATE', payload: path })
  },
  views: {
    fortune: {
      generate: () => ({ action: 'VIEWS.FORTUNE.GENERATE', payload: null })
    },
    catfact: {
      fetch: () => ({action: 'VIEWS.CATFACT.FETCH', payload: null})
    }
  }
}
export default actions
