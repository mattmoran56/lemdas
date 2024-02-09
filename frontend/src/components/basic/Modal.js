import React from "react";
import M from "react-modal";

const Modal = ({ children, isOpen }) => {
	return (
		<M
			isOpen={isOpen}
			style={{
				overlay: {
					backgroundColor: "rgba(0, 0, 0, 0.5)",
				},
				content: {
					top: "50%",
					left: "50%",
					right: "auto",
					bottom: "auto",
					marginRight: "-50%",
					transform: "translate(-50%, -50%)",
					width: "min-content",
				},
			}}
		>
			{children}
		</M>
	);
};

export default Modal;
