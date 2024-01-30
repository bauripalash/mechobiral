from collections import OrderedDict


RAW_FILE = "AkademiWorldList_raw.txt"


### Processing Raw Word List

def process_raw(doc : str, writefile : bool = False, output : str = "") -> None:
    with open(doc) as f:
        raw_text = f.read()
        lines = raw_text.splitlines()
        y = list(filter(lambda x: x != "", lines))

        if writefile:
            with open(output , "w") as w:
                text = "\n".join(y)
                w.write(text)

def preop_check(w : str) -> bool:
    return (w != "কি০"
            and w != "লি"
            and w != "কি"
            and w != "কী"
            and w != "তু০"
            and w != "যে০"
            and not w.startswith("-")
            and not w.startswith("+")
            and not w.startswith("<")
            and not w.startswith("ত্ব"))

def preop(input_file : str = "AkademiWorldList_Raw_Processed.txt") -> None:
    with open(input_file) as f:
        raw_text = f.read()
        lines = raw_text.splitlines()
        #print(lines)
        temp_list = list(filter(preop_check , lines))
        print(temp_list)
        result : list[str] = []

        for w in temp_list:
            result.extend(preop_indv(w))

        nodup_result = list(OrderedDict.fromkeys(result))


        with open("AkademiBanan_WorldList.txt" , "w") as wf:
            text = "\n".join(nodup_result)
            wf.write(text)

def preop_indv(word : str) -> list[str]:
    items = word.split(",")
    result = list(filter(preop_check, items))
    return result


def get_words(input_file : str) -> list[str]:
    with open(input_file) as f:
        raw_text = f.read()
        lines = raw_text.splitlines()
        return lines
        
def get_proc_word(w : str) -> list[str]:
    words = list(filter(lambda x: not x.startswith("(") and not x.startswith(",") and not x.startswith("'") , w.split(" ")))
    if len(words) < 1:
        return []
    result : list[str] = []
    for wd in words:
        if ( len(wd) > 1
                and not wd.startswith("+")
                and not wd.startswith("[") 
                and not wd.endswith("]")
                and not wd.startswith("(")
                and not wd.endswith(")")
                and not wd.startswith("'")
                and not wd.startswith("‘")
                and not wd.endswith("’")):
            result.append(wd)

    return result

def process_indv(words : list[str] = get_words("AkademiWorldList_Prep.txt")) -> None:
    words = get_words("AkademiWorldList_Prep.txt")
    result : list[str] = []
    for word in words:
        res = get_proc_word(word)
        if len(res) < 1:
            continue
        else:
            result.append(res[0].strip(","))

    with open("AkademiWorldList_Raw_Processed.txt" , "w") as f:
        f.write("\n".join(result))

def main_process_raw():
    process_raw(RAW_FILE , True , "AkademiWorldList_Prep.txt")
    process_indv()

### End of Processing Raw Words


if __name__ == "__main__":
    #main_process_raw()
    preop()

