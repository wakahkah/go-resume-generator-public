package service

import (
	"html/template"
	"strconv"

	"github.com/wakahkah/go-resume-generator/model"
	"github.com/wakahkah/go-resume-generator/utils"
)

func GetResumePDF(sourcePath string, outputPath string, templatePath string) (bool, error) {

	r := utils.NewRequestPdf("")

	resumeData, err := utils.GetDataFromFile(sourcePath)
	if err != nil {
		return false, err
	}

	ProcessData(resumeData)

	if err := r.ParseTemplate(templatePath, resumeData); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		return ok, nil
	} else {
		return false, err
	}
}

func ProcessData(resumeData *model.ResumeData) {
	// get skill string
	for gi := range resumeData.Skills {
		skillGroup := resumeData.Skills[gi]

		var skillNameAry []string
		for si := range skillGroup.SkillData {
			skillNameAry = append(skillNameAry, skillGroup.SkillData[si].Name)
		}

		resumeData.Skills[gi].SkillStr = utils.ArrayToStr(skillNameAry)
	}

	// get work exp date range string
	for i := range resumeData.WorkExperience {
		workExp := resumeData.WorkExperience[i]

		dateRangeStr := workExp.StartDate.Format("Jan ") + strconv.Itoa(workExp.StartDate.Year()) + " - "

		if workExp.IsPresent {
			dateRangeStr = dateRangeStr + "Present"
		} else {
			dateRangeStr = dateRangeStr + workExp.EndDate.Format("Jan ") + strconv.Itoa(workExp.EndDate.Year())
		}

		//achivement
		for achivedIndex := range workExp.Achieved {
			resumeData.WorkExperience[i].AchievedHtml = append(resumeData.WorkExperience[i].AchievedHtml, template.HTML(workExp.Achieved[achivedIndex]))
		}

		resumeData.WorkExperience[i].DateRangeStr = dateRangeStr
	}

	// get cert year
	for i := range resumeData.Certificates {
		cert := resumeData.Certificates[i]

		yearStr := strconv.Itoa(cert.Date.Year())

		resumeData.Certificates[i].YearStr = yearStr
	}

	// get award year
	for i := range resumeData.Awards {
		award := resumeData.Awards[i]

		yearStr := strconv.Itoa(award.Date.Year())

		resumeData.Awards[i].YearStr = yearStr
	}

	// get education date range string
	for i := range resumeData.Education {
		education := resumeData.Education[i]

		dateRangeStr := education.StartDate.Format("Jan ") + strconv.Itoa(education.StartDate.Year()) + " - "

		if education.IsPresent {
			dateRangeStr = dateRangeStr + "Present"
		} else {
			dateRangeStr = dateRangeStr + education.EndDate.Format("Jan ") + strconv.Itoa(education.EndDate.Year())
		}

		resumeData.Education[i].DateRangeStr = dateRangeStr
	}

	// get personal project range string
	for i := range resumeData.PersonalProject {
		pproj := resumeData.PersonalProject[i]

		dateRangeStr := pproj.StartDate.Format("Jan ") + strconv.Itoa(pproj.StartDate.Year())

		if pproj.IsPresent {
			dateRangeStr = dateRangeStr + " - Present"
		} else {
			if !pproj.EndDate.IsZero() {
				dateRangeStr = dateRangeStr + " - " + pproj.EndDate.Format("Jan ") + strconv.Itoa(pproj.EndDate.Year())
			}
		}

		resumeData.PersonalProject[i].DateRangeStr = dateRangeStr
	}

}
