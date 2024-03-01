import React, { useEffect, useState } from "react";

import Dataset from "../basic/Dataset";
import ErrorToast from "../../helpers/toast/errorToast";
import getStaredDatasets from "../../helpers/api/webApi/dataset/getStaredDatasets";

const StaredDatasets = () => {
	const [datasets, setDatasets] = useState([]);

	const getRecentDatasets = () => {
		getStaredDatasets().then((data) => {
			setDatasets(data);
		}).catch((error) => {
			ErrorToast(error);
		});
	};

	useEffect(() => {
		getRecentDatasets();
	}, []);

	return (
		<div className="mb-16">
			<h1 className="font-semibold text-2xl mx-2">Stared Datasets</h1>
			<div className="w-full h-[2px] mx-2 mb-4 bg-oxfordblue" />
			<div className="flex flex-wrap">
				{datasets.length === 0
					? (
						<p className="ml-2">
							Your stared datasets will appear here!
						</p>
					)
					: null}
				{datasets.map((dataset) => {
					return (
						<Dataset key={dataset.id} dataset={dataset} stared />
					);
				})}
			</div>
		</div>
	);
};

export default StaredDatasets;
