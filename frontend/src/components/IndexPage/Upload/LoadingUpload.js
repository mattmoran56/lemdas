import React from "react";
import Loader from "../../basic/Loader";

const LoadingUpload = () => {
	return (
		<div
			className={`text-oxfordblue border-oxfordblue w-full h-full rounded-md p-4 flex
											flex-col justify-center items-center border-2`}
		>
			<h1 className="text-2xl font-bold">Uploading...</h1>
			<Loader />
			<p className="text-xl">
				Please wait while we upload your file.
			</p>
		</div>
	);
};

export default LoadingUpload;
