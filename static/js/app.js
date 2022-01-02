const Prompt = () => {
  const toast = (c) => {
    const { msg = "", icon = "success", position = "top-end" } = c;

    const Toast = Swal.mixin({
      toast: true,
      title: msg,
      position,
      icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener("mouseenter", Swal.stopTimer);
        toast.addEventListener("mouseleave", Swal.resumeTimer);
      },
    });

    Toast.fire({});
  };

  const success = (c) => {
    const { msg = "", title = "", footer = "" } = c;
    console.log("clicked button and got success");
    Swal.fire({
      icon: "success",
      title,
      text: msg,
      footer,
    });
  };

  const error = (c) => {
    const { msg = "", title = "", footer = "" } = c;
    console.log("clicked button and got success");
    Swal.fire({
      icon: "error",
      title,
      text: msg,
      footer,
    });
  };

  const custom = async (c) => {
    const { icon = "", msg = "", title = "", showConfirmButton = true } = c;

    const { value: result } = await Swal.fire({
      icon,
      title,
      html: msg,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      showConfirmButton,
      willOpen: () => {
        if (c.willOpen !== undefined) {
          c.willOpen();
        }
      },
      didOpen: () => {
        if (c.didOpen !== undefined) {
          c.didOpen();
        }
      },
      preConfirm: () => {
        if (c.preConfirm !== undefined) {
          c.preConfirm();
        }
      },
    });

    if (result) {
      if (result.dismiss !== Swal.DismissReason.cancel) {
        if (result.value !== "") {
          if (c.callback !== undefined) {
            c.callback(result);
          }
        } else {
          c.callback(false);
        }
      } else {
        c.callback(false);
      }
    }
  };

  return {
    toast,
    success,
    error,
    custom,
  };
};
