class Elaparse {
  constructor(rootEl) {
    this.rootEl = rootEl
    this.formData = {
      webpage: '',
      selector: ''
    },
    this.status = ''
  }

  init() {
    this.handleStatus('pending')
    this.handleEvents()
  }

  async sendData() {
    this.handleStatus('loading')

    try {
      const cfg = {
          method: "POST",
        body: JSON.stringify({ webpage: this.formData.webpage, selector: this.formData.selector }),
        headers: {
          "Content-Type": "application/json",
        }
      }

      const response = await fetch('/api', cfg)
      const data = await response.json()

      if (response.ok && data) {
        this.handleStatus(data.message)
        this.handleData(data.result)
      } 
    } catch (err) {
      console.error(err)
      
      this.handleStatus('something went wrong, please, check your requirements again')
    }
  }

  handleData(data) {
    const tableEl = this.rootEl.querySelector('table')
    const tableHead = tableEl.querySelector('thead tr')
    const tableBody = tableEl.querySelector('tbody')

    if (!data) return

    tableHead.innerHTML = "<th>Result</th>"
    tableBody.innerHTML = data.reduce((prev, current) =>`${prev}<tr><td>${current}</td></tr>`, '')
  }

  handleStatus(message) {
    document.querySelector('form p span').textContent = this.status = message
  }

  handleEvents() {
    window.addEventListener('submit', e => {
      e.preventDefault()

      if (this.formData.webpage.trim() === '' || this.formData.selector.trim() === '') {
        this.handleStatus('webpage or selector fields are empty')
        return
      }

      this.sendData()
    })

    window.addEventListener('input', e => {
      const { target } = e;

      this.formData[target.id] = target.value
    })

    window.addEventListener('click', async e => {
      const { target } = e;
    })
  }
}

window.addEventListener('DOMContentLoaded', () => {
  window.$app = new Elaparse(document.querySelector('#app'))

  window.$app.init()
})

