import React from "react";

const Loader = () => {
	return (
		<div>
			<div
				className={`animate-spin-slow rounded-full p-2 bg-gradient-to-b
							from-oxfordblue-extralight to-oxfordblue m-2`}
			>
				<div className="w-10 h-10 bg-offwhite rounded-full" />
			</div>
		</div>
	);
};

export default Loader;
