import React from "react";

import RecentDatasets from "../../components/IndexPage/RecentDatasets";
import Upload from "../../components/IndexPage/Upload/index";

const LoggedInContent = () => {
	return (
		<div className="w-screen h-full bg-offwhite p-4 flex justify-center">

			<div className="flex max-w-7xl w-full">
				<div className="w-1/2 px-2">
					<RecentDatasets />
				</div>
				<div
					className="w-1/2 px-2 min-h-64 flex justify-center items-center"
				>
					<Upload />
				</div>
			</div>

		</div>
	);
};

export default LoggedInContent;
