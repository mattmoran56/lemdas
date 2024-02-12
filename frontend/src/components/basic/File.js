import React from "react";
import { PhotoIcon } from "@heroicons/react/24/outline";

const File = ({ file }) => {
	return (
		<div>
			<a
				className={`min-w-16 p-6 border-2 border-oxfordblue-200 m-2 flex flex-col items-center
						rounded-md shadow-md hover:shadow-lg
						hover:bg-gray-100 transition-colors duration-300`}
				href={`/file/${file.id}`}
			>
				<PhotoIcon className="h-8 w-8 mb-4" />
				<p className="font-semibold text-base">{file.name}</p>
			</a>
		</div>
	);
};

export default File;
