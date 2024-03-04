import React, { useEffect, useState } from "react";
import { PlusIcon, TrashIcon } from "@heroicons/react/24/outline";

import Button from "../basic/Button";
import getFileAttributes from "../../helpers/api/search/getFileAttributes";
import ErrorToast from "../../helpers/toast/errorToast";
import getDatasetAttributes from "../../helpers/api/search/getDatasetAttributes";

const OptionSection = ({
	first, lastRemaining, queryItem, onChange, onNewOption, onDelete,
}) => {
	const [fileAttributes, setFileAttributes] = useState([]);
	const [datasetAttributes, setDatasetAttributes] = useState([]);

	const [operand, setOperand] = useState("AND");
	const [field, setField] = useState("");
	const [value, setValue] = useState("");
	const [object, setObject] = useState("file");

	useEffect(() => {
		getFileAttributes().then((response) => {
			setFileAttributes(response.attributes);
		}).catch((error) => {
			ErrorToast(error);
		});

		getDatasetAttributes().then((response) => {
			setDatasetAttributes(response.attributes);
		}).catch((error) => {
			ErrorToast(error);
		});

		setOperand(queryItem.operand ?? "AND");
		setField(queryItem.field);
		setValue(queryItem.value);
		setObject(queryItem.object);
	}, []);

	useEffect(() => {
		onChange({
			operand, field, value, object,
		});
	}, [operand, field, value, object]);

	return (
		<div className="flex items-center justify-center my-2">
			<div className="w-24">
				{first ? null : (
					<select
						id="logic"
						className="text-black px-3 py-2 rounded-3xl border-2 mx-2"
						value={operand}
						onChange={(e) => { setOperand(e.target.value); }}
					>
						<option value="AND">AND</option>
						<option value="OR">OR</option>
						<option value="NOT">NOT</option>
					</select>
				)}
			</div>
			{/* TODO: Fit this to metadata (pull from server) */}
			<select
				id="field"
				className="text-black px-3 py-2 rounded-3xl border-2 mx-2 max-w-48 overflow-hidden"
				value={field}
				onChange={(e) => {
					console.log(e.target.value.split("-/--/-"));
					setField(e.target.value);
					setObject(e.target.value.split("-/--/-")[0]);
				}}
			>
				<option value="" disabled>File Attributes</option>
				<hr />
				{fileAttributes.map((attribute) => {
					return (
						<option
							key={attribute}
							value={`file-/--/-${attribute}`}
						>
							{attribute}
						</option>
					);
				})}

				<option value="" disabled>Dataset Attributes</option>
				<hr />
				{datasetAttributes.map((attribute) => {
					return (
						<option
							key={attribute}
							value={attribute}
							onSelect={() => { setObject("dataset"); }}
						>
							{attribute}
						</option>
					);
				})}
			</select>
			{/* TODO: Fit this to datatype (e.g. between for dates, contains for text) */}
			<p>=</p>
			<input
				className=" bg-transparent border-black border-b-2 mx-4 outline-none"
				placeholder="Search term"
				value={value}
				id="term"
				onChange={(e) => { setValue(e.target.value); }}
				autoComplete="off"
			/>
			<Button className="mx-2 rounded-full px-2" onClick={onNewOption}>
				<PlusIcon className="w-4 h-4" />
			</Button>
			{lastRemaining ? <div className="w-4 mx-4" /> : (
				<Button className="mx-2 rounded-full px-2 bg-oxfordblue border-oxfordblue" onClick={onDelete}>
					<TrashIcon className="w-4 h-4" />
				</Button>
			)}
		</div>
	);
};

export default OptionSection;
