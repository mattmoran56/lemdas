import React from "react";
import { PhotoIcon } from "@heroicons/react/24/outline";
import Loader from "./Loader";

const File = ({ file }) => {
	return (
		file.status === "support_processed"
			? null
			: (
				<div>
					<a
						className={`w-44 h-32 p-6 border-2 border-oxfordblue-200 m-2 flex flex-col items-center
								rounded-md shadow-md hover:shadow-lg justify-center
								hover:bg-gray-100 transition-colors duration-300
								${file.status !== "processed" ? "bg-gray-200" : "bg-white"}`}
						href={`/file/${file.id}`}
					>
						{file.status === "processed"
							? <PhotoIcon className="h-8 w-8 mb-4" />
							: <Loader className="h-6 w-6" outerClassName="p-1.5 mt-0" />}
						<p className="font-semibold text-base">
							{file.name.length > 10
								? `${file.name.substring(0, 10)}...`
								: file.name}
						</p>
					</a>
				</div>
			)
	);
};

export default File;
