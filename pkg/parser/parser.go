package parser

type Object struct {
	Name       string
	Objects    map[string][]Object
	Attributes map[string][]Attribute
}

func Parse(lines []string) ([]Object, error) {
	structureElements, err := readStructure(lines)
	if err != nil {
		return nil, err
	}

	objects, err := convertStructureElements(structureElements)
	if err != nil {
		return nil, err
	}

	result := make([]Object, 0, len(structureElements))

	for _, v := range objects {
		result = append(result, v...)
	}

	return result, nil
}

func (element structureElement) toObject() (*Object, error) {
	objects, err := convertStructureElements(element.structureElements)
	if err != nil {
		return nil, err
	}

	attributes, err := convertAttributes(element.attributes)
	if err != nil {
		return nil, err
	}

	return &Object{
		Name:       element.name,
		Objects:    objects,
		Attributes: attributes,
	}, nil
}

func convertStructureElements(elements []*structureElement) (map[string][]Object, error) {
	objects := make(map[string][]Object)

	for _, subElement := range elements {
		object, err := subElement.toObject()
		if err != nil {
			return nil, err
		}

		previous := objects[object.Name]
		previous = append(previous, *object)
		objects[object.Name] = previous
	}

	return objects, nil
}

func convertAttributes(attributeLines []string) (map[string][]Attribute, error) {
	attributes := make(map[string][]Attribute, len(attributeLines))

	for _, attrLine := range attributeLines {
		attr, err := readAttributeLine(attrLine)
		if err != nil {
			return nil, err
		}

		previous := attributes[attr.Name]
		previous = append(previous, *attr)
		attributes[attr.Name] = previous
	}

	return attributes, nil
}
