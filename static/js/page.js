v = new Vue({
  el: '#main',
  beforeMount() {
    window.addEventListener('keydown', handleKeydown, null);
  },
  beforeDestroy() {
    window.removeEventListener('keydown', handleKeydown);
  }
})
