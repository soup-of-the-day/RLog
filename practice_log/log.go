package practice_log

// Holds data about an individual rudiment
type PracticeEntry struct {
	BPM    int    // BPM to practice this rudiment at (the user's previous best!)
	Name   string // Full name of the rudiment
}

// List of PracticeEntries + behavior for picking random entries as well as updating records
type PracticeLog struct {
	Log map[string]*PracticeEntry // map of entries. 3-letter Abbreviation -> PracticeEntry
}

func NewPracticeLog(entriesMap map[string]*PracticeEntry) *PracticeLog {
	return &PracticeLog{
		Log: entriesMap,
	}
}

// Get random entries from the list for practice
func (p *PracticeLog) GetRandomEntries(numEntries int) []*PracticeEntry {
	entries := make([]*PracticeEntry, 0)
	for _, entry := range p.Log { // ranging over a map returns elements in some random order
		entries = append(entries, entry)
	}
	return entries[0:numEntries]
}

// Update the entry in the list associated with that abbrev
func (p *PracticeLog) UpdateEntry(abbrev string, bpm int) {
	entry := p.Log[abbrev]
	entry.BPM = bpm
}
