package console

import log "github.com/sirupsen/logrus"

func (u *UserInterface) createConsoleWorker() {
	body, err := u.httpRepo.GetRequestBody(u.startUrl)
	if err != nil {
		log.Errorf("console.UserInterface.createConsoleWorker()->GetRequestBody: %v", err)
	}
	resultRender, err := u.uc.RenderBodyToAscii(string(body))
	if err != nil {
		log.Errorf("console.UserInterface.createConsoleWorker()->RenderBodyToAscii: %v", err)
	}
	log.Info(resultRender)
}
