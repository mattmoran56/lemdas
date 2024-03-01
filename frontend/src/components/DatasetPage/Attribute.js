import React, { useEffect, useState } from "react";
import { PlusIcon, TrashIcon } from "@heroicons/react/24/outline";

import addAttribute from "../../helpers/api/webApi/datasetAttributes/addAttribute";
import deleteDatasetAttribute from "../../helpers/api/webApi/datasetAttributes/deleteDatasetAttribute";
import updateDatasetAttribute from "../../helpers/api/webApi/datasetAttributes/updateDatasetAttribute";
import ErrorToast from "../../helpers/toast/errorToast";

const Attribute = ({
	attribute, datasetId, setNeedRefresh, addNewAttribute, writeAccess,
}) => {
	const [clicked, setClicked] = useState(false);
	const [error, setError] = useState(false);

	const [attributeId, setAttributeId] = useState("");
	const [attributeName, setAttributeName] = useState("");
	const [value, setValue] = useState("");

	const handleClickOff = () => {
		setClicked(false);
		if (attributeName === "" || value === "") {
			setError(true);
		} else if (attribute.id === undefined) {
			addAttribute(datasetId, attributeName, value).then((d) => {
				addNewAttribute();
				setAttributeId(d.id);
			}).catch((e) => {
				ErrorToast(e);
			});
		} else {
			updateDatasetAttribute(datasetId, attribute.id, attributeName, value).catch((e) => {
				ErrorToast(e);
			});
		}
	};

	const handleDelete = () => {
		deleteDatasetAttribute(datasetId, attribute.id).then(() => {
			setNeedRefresh(true);
		}).catch((e) => {
			ErrorToast(e);
		});
	};

	useEffect(() => {
		setAttributeId(attribute.id);
		setAttributeName(attribute.attribute_name);
		setValue(attribute.attribute_value);
	}, []);

	useEffect(() => {
		if (attributeName !== "" && value !== "") {
			setError(false);
		}
	}, [attributeName, value]);

	return (
		<tr>
			<td className="text-right font-light pr-2 flex" aria-label="New attribute">
				<input
					className={`outline-none border-b-2 w-full text-right pl-2 mt-[2px] border-transparent
								${error ? "border-red-500 bg-red-100" : ""}
								${attributeName === "" || value === "" || clicked ? "border-gray-300" : ""}`}
					placeholder="New attribute"
					onFocus={() => { setClicked(true); }}
					onBlur={handleClickOff}
					value={attributeName}
					onChange={(e) => { setAttributeName(e.target.value); }}
					disabled={!writeAccess}
				/>:
			</td>
			<td className="font-medium" aria-label="New Value">
				<input
					className={`outline-none border-b-2 flex-auto px-2 border-transparent
								${error ? "border-red-500 bg-red-100" : ""}
								${attributeName === "" || value === "" || clicked ? "border-gray-300" : ""}`}
					placeholder="value"
					onFocus={() => { setClicked(true); }}
					onBlur={handleClickOff}
					value={value}
					onChange={(e) => { setValue(e.target.value); }}
					disabled={!writeAccess}
				/>
			</td>
			<td>
				{attributeId
					? (
						<button
							className={writeAccess ? "" : "hidden"}
							type="button"
							aria-label="delete"
							onClick={handleDelete}
							disabled={!writeAccess}
						>
							<TrashIcon className="h-6 w-6 text-gray-500 hover:text-red-500" />
						</button>
					)
					: (
						<button
							className={writeAccess ? "" : "hidden"}
							type="button"
							aria-label="save"
							disabled={!writeAccess}
						>
							<PlusIcon className="h-6 w-6 text-gray-300 hover:text-gray-500" />
						</button>
					)}
			</td>
		</tr>
	);
};

const Attributes = ({
	attributes, datasetId, setNeedRefresh, writeAccess,
}) => {
	const [attributeList, setAttributeList] = useState([]);

	const addNewAttribute = () => {
		setAttributeList([...attributeList, { attribute_name: "", attribute_value: "" }]);
	};

	useEffect(() => {
		setAttributeList([...attributes, {
			attribute_name: "",
			attribute_value: "",
		}]);
	}, [attributes]);

	return (
		<div className="w-full">
			<table className="w-fit mb-4">
				<tbody>
					{attributeList.map((attribute) => {
						return (
							<Attribute
								key={attribute.id ? attribute.id : "new"}
								attribute={attribute}
								datasetId={datasetId}
								setNeedRefresh={setNeedRefresh}
								addNewAttribute={addNewAttribute}
								writeAccess={writeAccess}
							/>
						);
					})}
				</tbody>
			</table>
		</div>
	);
};

export default Attributes;
