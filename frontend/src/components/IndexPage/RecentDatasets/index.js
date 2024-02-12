import React, { useEffect, useState } from "react";
import { FolderPlusIcon } from "@heroicons/react/24/outline";

import Dataset from "./Dataset";
import NewDatasetModal from "./NewDatasetModal";
import getDatasets from "../../../helpers/api/webApi/dataset/getDatasets";

const RecentDatasets = () => {
	const [datasets, setDatasets] = useState([]);
	const [showNewDatasetModal, setShowNewDatasetModal] = useState(false);

	const handleNewDataset = () => {
		setShowNewDatasetModal(true);
	};

	const getRecentDatasets = () => {
		getDatasets("updated_at").then((data) => {
			setDatasets(data);
		});
	};

	useEffect(() => {
		getRecentDatasets();
	}, []);

	return (
		<div>
			<NewDatasetModal
				isOpen={showNewDatasetModal}
				setIsOpen={setShowNewDatasetModal}
				onClose={getRecentDatasets}
			/>
			<h1 className="font-semibold text-2xl mx-2">Recent Datasets</h1>
			<div className="w-full h-[2px] mx-2 mb-4 bg-oxfordblue" />
			<div className="flex flex-wrap">
				{datasets.map((dataset) => {
					return (
						<Dataset key={dataset.id} dataset={dataset} />
					);
				})}
				<button
					className={`w-[calc(50%-1rem)] p-2 border-2 border-oxfordblue-200 m-2 flex items-center rounded-md
						shadow-md border-dashed
						hover:bg-gray-100 transition-colors duration-300`}
					onClick={handleNewDataset}
					type="button"
				>
					<FolderPlusIcon className="h-8 w-8 mr-2" />
					<p className="font-semibold text-base">New Dataset</p>
				</button>
			</div>
		</div>
	);
};

export default RecentDatasets;
