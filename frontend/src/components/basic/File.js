import React, { useState, useEffect } from "react";
import { ExclamationTriangleIcon, PhotoIcon } from "@heroicons/react/24/outline";
import Loader from "./Loader";
import getPreviewURL from "../../helpers/api/webApi/file/getPreview";

const File = ({ file }) => {
	const [previewUrl, setPreviewUrl] = useState("");

	useEffect(() => {
		if (file.status === "processed") {
			getPreviewURL(file.id)
				.then((data) => {
					setPreviewUrl(data.url);
				});
		}
	}, []);

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
						{file.status === "uploaded" || file.status === "processing"
							? <Loader className="h-6 w-6" outerClassName="p-1.5 mt-0" />
							: null}
						{file.status === "awaitingtxt"
							? <ExclamationTriangleIcon className="h-6 w-6 text-indianred" />
							: null}
						{file.status === "processed" && previewUrl !== ""
							? <img src={previewUrl} alt="preview" className="h-12 w-auto mb-2" />
							: null}
						{file.status === "processed" && previewUrl === ""
							? <PhotoIcon className="h-6 w-6 text-oxfordblue-400" />
							: null}
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
