package axmlParser

type AndroidListener struct {
	AppName          string
	Icon             string
	PackageName      string
	VersionName      string
	VersionCode      string
	ActivityName     string
	tempActivityName string
	findMainActivity bool
}

func (listener *AndroidListener) StartDocument() {
}

/**
 * Receive notification of the end of a document.
 */
func (listener *AndroidListener) EndDocument() {
}

/**
 * Begin the scope of a prefix-URI Namespace mapping.
 *
 * @param prefix
 *            the Namespace prefix being declared. An empty string is used
 *            for the default element namespace, which has no prefix.
 * @param uri
 *            the Namespace URI the prefix is mapped to
 */
func (listener *AndroidListener) StartPrefixMapping(prefix, uri string) {
}

/**
 * End the scope of a prefix-URI mapping.
 *
 * @param prefix
 *            the prefix that was being mapped. This is the empty string
 *            when a default mapping scope ends.
 * @param uri
 *            the Namespace URI the prefix is mapped to
 */
func (listener *AndroidListener) EndPrefixMapping(prefix, uri string) {}

/**
 * Receive notification of the beginning of an element.
 *
 * @param uri
 *            the Namespace URI, or the empty string if the element has no
 *            Namespace URI or if Namespace processing is not being
 *            performed
 * @param localName
 *            the local name (without prefix), or the empty string if
 *            Namespace processing is not being performed
 * @param qName
 *            the qualified name (with prefix), or the empty string if
 *            qualified names are not available
 * @param atts
 *            the attributes attached to the element. If there are no
 *            attributes, it shall be an empty Attributes object. The value
 *            of this object after startElement returns is undefined
 */
func (listener *AndroidListener) StartElement(uri, localName, qName string,
	attrs []*Attribute) {
	if listener.findMainActivity {
		return
	}

	if localName == "manifest" {
		for _, attr := range attrs {
			switch attr.Name {
			case "package":
				listener.PackageName = attr.Value
				break
			case "versionCode":
				listener.VersionCode = attr.Value
				break
			case "versionName":
				listener.VersionName = attr.Value
				break
			}
		}
		return
	}

	if localName == "application" {
		for _, attr := range attrs {
			switch attr.Name {
			case "label":
				listener.AppName = attr.Value
				break
			case "icon":
				listener.Icon = attr.Value
				break
			}
		}
		return
	}

	if localName == "activity" {
		for _, attr := range attrs {
			if attr.Name == "name" && attr.Prefix == "android" &&
				attr.Namespace == "http://schemas.android.com/apk/res/android" {
				listener.tempActivityName = attr.Value
				break
			}
		}
	}

	if localName != "action" {
		return
	}

	//fmt.Println(uri, localName, qName)
	for _, attr := range attrs {
		if attr.Name == "name" && attr.Prefix == "android" &&
			attr.Namespace == "http://schemas.android.com/apk/res/android" &&
			attr.Value == "android.intent.action.MAIN" {
			listener.ActivityName = listener.tempActivityName
			listener.findMainActivity = true
			break
		}
	}
}

/**
 * Receive notification of the end of an element.
 *
 * @param uri
 *            the Namespace URI, or the empty string if the element has no
 *            Namespace URI or if Namespace processing is not being
 *            performed
 * @param localName
 *            the local name (without prefix), or the empty string if
 *            Namespace processing is not being performed
 * @param qName
 *            the qualified XML name (with prefix), or the empty string if
 *            qualified names are not available
 */
func (listener *AndroidListener) EndElement(uri, localName, qName string) {}

/**
 * Receive notification of text.
 *
 * @param data
 *            the text data
 */
func (listener *AndroidListener) Text(data string) {}

/**
 * Receive notification of character data (in a <![CDATA[ ]]> block).
 *
 * @param data
 *            the text data
 */
func (listener *AndroidListener) CharacterData(data string) {}

/**
 * Receive notification of a processing instruction.
 *
 * @param target
 *            the processing instruction target
 * @param data
 *            the processing instruction data, or null if none was supplied.
 *            The data does not include any whitespace separating it from
 *            the target
 * @throws org.xml.sax.SAXException
 *             any SAX exception, possibly wrapping another exception
 */
func (listener *AndroidListener) ProcessingInstruction(target, data string) {

}
