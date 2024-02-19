import React from "react";
import { FolderIcon } from "@heroicons/react/24/outline";

const Dataset = ({ dataset }) => {
	return (
		<a
			className="border-[1px] border-gray-500 rounded-md flex p-4 my-2 shadow-md hover:underline"
			href={`/dataset/${dataset.id}`}
		>
			<div className="mr-8">
				<FolderIcon className="h-10 w-10" />
			</div>
			<div>
				<h2 className="font-semibold text-xl">{dataset.dataset_name}</h2>
				<div className="w-full flex !no-underline">
					<p className="text-gray-800 mr-4 !no-underline">Author: </p>
					<p className="font-medium !no-underline">{dataset.owner_name}</p>
				</div>
			</div>
		</a>
	);
};

export default Dataset;
