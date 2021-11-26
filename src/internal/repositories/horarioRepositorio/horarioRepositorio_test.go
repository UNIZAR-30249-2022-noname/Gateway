package horarioRepositorio_test

import (
	"testing"

	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/internal/repositories/horarioRepositorio"
	"github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/apperrors"
	consultas "github.com/D-D-EINA-Calendar/CalendarServer/src/pkg/sql"
	"github.com/stretchr/testify/assert"
)


func TestGetAvaiableHours(t *testing.T) {
	
	//Prepare
	assert := assert.New(t)
	hoursexpected := []domain.AvailableHours{
		{
			Subject:   domain.Subject{Kind: 1,Name: "Proyecto Software"},
			Remaining: 30,
			Max:       30,
		},
		{
			Subject:   domain.Subject{Kind: 2, Name: "Sistemas Operativos"},
			Remaining: 25,
			Max:       25,
		},
	}
	ternaAsked := domain.Terna{
		Titulacion: "Ing. Informatica",
		Curso:      1,
		Grupo:      "1",
	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)

	//Start
	hoursgot, _ := repos.GetAvailableHours(ternaAsked)

	assert.Equal(len(hoursgot), len(hoursexpected), "Should be the same length")
	for i, h := range hoursgot {
		assert.Equal(h, hoursexpected[i], "Should be the same AvaiableHours")
	}

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion)
	repos.CloseConn()
}

func TestCreateEntry(t *testing.T) {

	//Prepare
	assert := assert.New(t)
	entryAsked := domain.Entry{

		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",

	}
	
	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Aula1);

	//Start (Everything OK)
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(repos.EntryFound(entryAsked), true, "Should be the same entries")

	//Start (Empty name)
	entryAsked = domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: ""},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",
	}

	err := repos.CreateNewEntry(entryAsked)
	if err != nil { assert.Equal(apperrors.ErrSql, err, "Should be the same error") }

	//Start (Empty room)
	entryAsked = domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room: domain.Room{Name: ""},
		Week: "",
		Group: "",
	}

	err = repos.CreateNewEntry(entryAsked)
	if err != nil { assert.Equal(apperrors.ErrSql, err, "Should be the same error") }

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

//Create entry practicas & try to create entry practicas without the week and group
func TestCreateEntryPract(t *testing.T) {

	//Prepare
	assert := assert.New(t)
	entryAsked := domain.Entry{

		Init: domain.NewHour(2,50),
		End: domain.NewHour(4,50),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room: domain.Room{Name: "2"},
		Week: "a",
		Group: "mananas",

	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Hora12)
	repos.RawExec(consultas.Aula1);			repos.RawExec(consultas.Aula2)

	//Start
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(repos.EntryFound(entryAsked), true, "Should be the same entries")

	//Empty group
	entryAsked = domain.Entry{

		Init: domain.NewHour(5,30),
		End: domain.NewHour(6,20),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room: domain.Room{Name: "3"},
		Week: "a",
		Group: "",
	}

	err := repos.CreateNewEntry(entryAsked)
	if err != nil { assert.Equal(apperrors.ErrInvalidKind, err, "Should be the same error") }

	//Empty group
	entryAsked = domain.Entry{

		Init: domain.NewHour(5,30),
		End: domain.NewHour(6,20),
		Subject: domain.Subject{Kind: 2, Name: "Proyecto Software"},
		Room: domain.Room{Name: "3"},
		Week: "",
		Group: "mananas",
	}

	err = repos.CreateNewEntry(entryAsked)
	if err != nil { assert.Equal(apperrors.ErrInvalidKind, err, "Should be the same error") }


	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

//Create entry problemas & try to create entry problemas without the group
func TestCreateEntryProb(t *testing.T) {
	assert := assert.New(t)
	entryAsked := domain.Entry{

		Init: domain.NewHour(5,30),
		End: domain.NewHour(6,20),
		Subject: domain.Subject{Kind: 3, Name: "Proyecto Software"},
		Room: domain.Room{Name: "3"},
		Week: "",
		Group: "niapar",

	}

	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Hora12);		repos.RawExec(consultas.Hora13);
	repos.RawExec(consultas.Aula1);			repos.RawExec(consultas.Aula2)
	repos.RawExec(consultas.Aula3)

	//Start (Everything Ok)
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(repos.EntryFound(entryAsked), true, "Should be the same entries")

	//Empty group
	entryAsked = domain.Entry{

		Init: domain.NewHour(5,30),
		End: domain.NewHour(6,20),
		Subject: domain.Subject{Kind: 3, Name: "Proyecto Software"},
		Room: domain.Room{Name: "3"},
		Week: "",
		Group: "",
	}

	err := repos.CreateNewEntry(entryAsked)
	if err != nil { assert.Equal(apperrors.ErrInvalidKind, err, "Should be the same error") }

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}

func TestDeleteEntry(t *testing.T) {
	//Prepare
	assert := assert.New(t)
	entryAsked := domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",
	}
	
	repos := horarioRepositorio.New()
	repos.RawExec(consultas.Titulacion1); 	repos.RawExec(consultas.Titulacion2)
	repos.RawExec(consultas.Curso1); 		repos.RawExec(consultas.Curso2)
	repos.RawExec(consultas.Asignatura1); 	repos.RawExec(consultas.Asignatura2)
	repos.RawExec(consultas.Grupodocente1); repos.RawExec(consultas.Grupodocente2)
	repos.RawExec(consultas.Hora1); 		repos.RawExec(consultas.Hora2)
	repos.RawExec(consultas.Aula1);

	//Start (Everything Ok)
	repos.CreateNewEntry(entryAsked)
	
	assert.Equal(repos.EntryFound(entryAsked), true, "Should have been deleted")

	repos.DeleteEntry(entryAsked)

	assert.Equal(repos.EntryFound(entryAsked), false, "")


	//Empty name
	entryAsked = domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: ""},
		Room: domain.Room{Name: "1"},
		Week: "",
		Group: "",
	}

	err := repos.DeleteEntry(entryAsked)
	if err != nil { assert.Equal(apperrors.ErrSql, err, "Should be the same error") }

	//Empty room
	entryAsked = domain.Entry{
		Init: domain.NewHour(1,30),
		End: domain.NewHour(2,40),
		Subject: domain.Subject{Kind: 1, Name: "Proyecto Software"},
		Room: domain.Room{Name: ""},
		Week: "",
		Group: "",
	}

	err = repos.DeleteEntry(entryAsked)
	if err != nil { assert.Equal(apperrors.ErrSql, err, "Should be the same error") }

	//Delete
	repos.RawExec(consultas.TruncHora); 		repos.RawExec(consultas.TruncGrupo)
	repos.RawExec(consultas.TruncAsignatura); 	repos.RawExec(consultas.TruncCurso)
	repos.RawExec(consultas.TruncTitulacion);	repos.RawExec(consultas.TruncAula)
	repos.RawExec(consultas.TruncEntry)
	repos.CloseConn()
}
