import React from "react";

const Loader = ({ className, outerClassName }) => {
	return (
		<div>
			<div
				className={`animate-spin-slow rounded-full p-2 bg-gradient-to-b
							from-oxfordblue-extralight to-oxfordblue m-2 ${outerClassName}`}
			>
				<div className={`w-10 h-10 bg-offwhite rounded-full ${className}`} />
			</div>
		</div>
	);
};

export default Loader;
