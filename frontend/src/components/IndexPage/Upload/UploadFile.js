import React from "react";
import { CloudArrowDownIcon } from "@heroicons/react/24/solid";
import { FileUploader } from "react-drag-drop-files";

const UploadFile = ({ setFile }) => {
	const handleChangeFile = (f) => {
		setFile(f);
	};

	return (
		<FileUploader
			handleChange={handleChangeFile}
			name="file"
			multiple
			enctype="multipart/form-data"
			classes={`text-oxfordblue !border-oxfordblue !w-full !h-full !rounded-md !p-4 !flex
													!justify-center !items-center !border-dashed !border-2`}
		>
			<div
				className={`text-oxfordblue text-center flex flex-col justify-center
													items-center`}
			>
				<CloudArrowDownIcon className="w-8 h-8" />
				Drag and drop a file here, or click to select a file to upload.
			</div>
		</FileUploader>
	);
};

export default UploadFile;
