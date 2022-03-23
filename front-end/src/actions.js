export default {
  session: {
    updateUsername: un => ({ action: 'SESSION.UPDATE_USERNAME', payload: un }),
    navigate: path => ({ action: 'SESSION.NAVIGATE', payload: path })
  }
}
