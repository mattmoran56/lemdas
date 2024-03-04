import React from "react";

import RecentDatasets from "../../components/IndexPage/RecentDatasets";
import Upload from "../../components/IndexPage/Upload/index";
import StaredDatasets from "../../components/IndexPage/StaredDatasets";
import Groups from "../../components/IndexPage/Groups";

const LoggedInContent = () => {
	return (
		<div className="w-screen h-full bg-offwhite p-4 flex justify-center">

			<div className="flex max-w-7xl w-full h-full">
				<div className="w-1/2 px-2">
					<StaredDatasets />
					<RecentDatasets />
				</div>
				<div
					className="w-1/2 px-2 min-h-96 flex justify-center items-center flex-col"
				>
					<Upload />
					<Groups />
				</div>
			</div>

		</div>
	);
};

export default LoggedInContent;
