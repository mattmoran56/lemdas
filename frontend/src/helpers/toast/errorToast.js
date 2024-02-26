import { toast } from "react-toastify";

const ErrorToast = (error) => {
	toast.error((`Error: ${error.message}`), {
		position: "top-right",
		autoClose: 5000,
		hideProgressBar: false,
		closeOnClick: true,
		progress: undefined,
		theme: "colored",
		className: "!bg-red mt-20",
	});
};

export default ErrorToast;
