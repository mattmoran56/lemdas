import React from "react";

import Index from "../../components/IndexPage/Upload";

const LoggedInContent = () => {
	return (
		<div className="w-screen h-full bg-offwhite p-4 flex justify-center">

			<div className="flex max-w-7xl w-full">
				<div className="w-1/2">
					<h1 className="text-2xl font-bold">Welcome back!</h1>
					<p className="text-xl">Youre logged in.</p>
				</div>
				<div
					className="w-1/2 min-h-64 flex justify-center items-center"
				>
					<Index />
				</div>
			</div>

		</div>
	);
};

export default LoggedInContent;
