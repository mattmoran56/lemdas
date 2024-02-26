import React from "react";
import { FolderIcon } from "@heroicons/react/24/outline";
import { StarIcon } from "@heroicons/react/24/solid";

const Dataset = ({ dataset, stared }) => {
	return (
		<a
			className={`w-[calc(50%-1rem)] p-2 border-2 border-oxfordblue-200 m-2 flex items-center rounded-md
						shadow-md
						hover:bg-gray-100 transition-colors duration-300`}
			href={`/dataset/${dataset.id}`}
		>
			<FolderIcon className="h-8 w-8 mr-2" />
			<p className="font-semibold text-base">{dataset.dataset_name}</p>
			<div className="flex-grow" />
			{stared
				? <StarIcon className="h-6 w-6 ml-2 text-yellow-500" />
				: null}
		</a>
	);
};

export default Dataset;
