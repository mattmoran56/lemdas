import React from "react";

const Button = ({ children, onClick, className }) => {
	return (
		<button
			className={`bg-indianred text-lightlavender px-4 py-2 rounded-3xl border-2 border-indianred
						flex items-center justify-center
						${className}`}
			type="button"
			onClick={onClick}
		>
			{children}
		</button>
	);
};

export default Button;
