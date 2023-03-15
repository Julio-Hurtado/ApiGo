package repository

import (
	"api/db"
	model "api/models"
)

func ObtenerDatosAsistencia(codigoAsistencia string) (model.Asistencia, error) {
	db.AbrirConexion()
	sqlQuery := "select * from fn_webservice_obtener_datos_asistencia($1)"
	asistencia := model.Asistencia{}
	err := db.Db.QueryRow(sqlQuery, codigoAsistencia).Scan(
		&asistencia.TipoAsistencia,
		&asistencia.CodigoAsistencia,
		&asistencia.FechaAsistencia,
		&asistencia.NumeroAfiliacion,
		&asistencia.FinanciadorId,
		&asistencia.SiglasFinanciador,
		&asistencia.NombreFinanciador,
		&asistencia.CodigoCargoFacturacion,
		&asistencia.NombreCargoFacturacion,
		&asistencia.TipoAsegurado,
		&asistencia.MedicoId,
		&asistencia.CodigoMedico,
		&asistencia.SexoMedico,
		&asistencia.PrimerNombreMedico,
		&asistencia.SegundoNombreMedico,
		&asistencia.PrimerApellidoMedico,
		&asistencia.SegundoApellidoMedico,
		&asistencia.NombreCompletoMmedico,
		&asistencia.PacienteId,
		&asistencia.NombrePaciente,
		&asistencia.ServicioId,
		&asistencia.CodigoServicio,
		&asistencia.NombreServicio,
		&asistencia.EspecialidadId,
		&asistencia.CodigoEspecialidad,
		&asistencia.NombreEspecialidad,
		&asistencia.CambiarServicio,
		&asistencia.NumeroExpediente,
		&asistencia.CodigoUnidadOrigen,
		&asistencia.NombreUnidadOrigen,
		&asistencia.CodigoAlternoEstablecimiento,
		&asistencia.NombreEstablecimiento,
	)

	if err != nil {
		return asistencia, err
	}
	db.CerrarConexion()

	return asistencia, nil

}
