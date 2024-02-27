import React from "react";

const Button = ({
	children, onClick, className, disabled,
}) => {
	return (
		<button
			className={`bg-indianred text-lightlavender px-4 py-2 rounded-3xl border-2 border-indianred
						flex items-center justify-center
						${disabled ? "bg-gray-400 !border-gray-400" : ""}
						${className}`}
			type="button"
			onClick={onClick}
			disabled={disabled}
		>
			{children}
		</button>
	);
};

export default Button;
