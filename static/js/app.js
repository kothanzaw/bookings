let attention = Prompt()

function notify(msg, msgType) {
  notie.alert({
    type: msgType, // optional, default = 4, enum: [1, 2, 3, 4, 5, 'success', 'warning', 'error', 'info', 'neutral']
    text: msg,
  })
}
function notifyModal(title, text, icon, confirmationButtonText) {
  Swal.fire({
    title: title,
    html: text,
    icon: icon,
    confirmButtonText: confirmationButtonText,
  })
}
function Prompt() {
  let toast = function (c) {
    const { msg = '', icon = 'success', position = 'top-end' } = c
    const Toast = Swal.mixin({
      toast: true,
      title: msg,
      position: position,
      icon: icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
      },
    })

    Toast.fire({})
  }

  let success = function (c) {
    const {
      msg = '',
      title = '',

      footer = '',
    } = c
    Swal.fire({
      icon: 'success',
      title: title,
      text: msg,
      footer: footer,
    })
  }
  let error = function (c) {
    const {
      msg = '',
      title = '',

      footer = '',
    } = c
    Swal.fire({
      icon: 'error',
      title: title,
      text: msg,
      footer: footer,
    })
  }
  async function custom(c) {
    const { msg = '', title = '' } = c

    const { value: result } = await Swal.fire({
      title: title,
      html: msg,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      willOpen: () => {
        if (c.willOpen !== undefined) {
          c.willOpen()
        }
      },
      didOpen: () => {
        if (c.didOpen != undefined) { 
          c.didOpen();
        }
      },
      preConfirm: () => {
        if (c.preConfirm !== undefined) { 
          c.preConfirm();
        }
      },
    })

    if (result) {
      // /Swal.fire(JSON.stringify(formValues))
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== '') {
          if (c.callback !== undefined) {
            c.callback(result)
          }
        } else {
          c.callback(false)
        }
      } else {
        c.callback(false)
      }
    }
  }

  return {
    toast,
    success,
    error,
    custom,
  }
}
