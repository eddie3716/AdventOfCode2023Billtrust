using System;
using System.Data.SqlClient;
using System.IO;
using System.CommandLine;
using System.CommandLine.Invocation;
using System.CommandLine.NamingConventionBinder;
using System.CommandLine.Parsing;

class Program
{
    static async Task<int> Main(string[] args)
    {
        var rootCommand = new RootCommand
        {
            new Option<string>(
                "--server",
                description: "The database server (required)"),
            new Option<string>(
                "--db",
                description: "The database (required)"),
            new Option<string>(
                "--envelopeid",
                description: "The envelope ID (optional...either --envelopeid or --enveloperange must be present)"),
            new Option<string>(
                "--enveloperange",
                description: "The envelope ID range (optional...either --envelopeid or --enveloperange must be present): format: start-end, example: 1-100"),
        };

        rootCommand.Description = "Exports 1 or more files from the AR database based on the envelope ID or envelope ID range.\n" +
                          "Uses integration authentication, so the user running this program must have read access to the database.\n\n" +
                          "Example usage:\n" +
                          "ExportEnvelopes --server cac-readonly.btazure.local --db MTProd_BtApiTest1 --envelopeid 123\n" +
                          "ExportEnvelopes --server cac-readonly.btazure.local --db MTProd_BtApiTest1 --enveloperange 1-100";
        rootCommand.Handler = CommandHandler.Create<string, string, string, string>((server, db, envelopeid, enveloperange) =>
        {
            if (string.IsNullOrEmpty(server))
            {
                Console.WriteLine("Server is required. Use the -h option for help.");
                return;
            }
            if (string.IsNullOrEmpty(db))
            {
                Console.WriteLine("Database is required. Use the -h option for help.");
                return;
            }
            if (string.IsNullOrEmpty(envelopeid) && string.IsNullOrEmpty(enveloperange))
            {
                Console.WriteLine("Either -envelopeid or -enveloperange must be present.  Use the -h option for help.");
                return;
            }
            var query = "";

            if (!string.IsNullOrEmpty(envelopeid))
            {
                if (!int.TryParse(envelopeid, out int id))
                {
                    Console.WriteLine("Envelope ID must be an integer.  Use the -h option for help.");
                    return;
                }
                query = $@"
                SELECT P.EnvelopeID, P.Sequence, FT.Description as Extension, F. Content 
                FROM [AR].[Page] P 
                INNER JOIN [AR].[PageFile] PF on PF.PageID = P.PageID 
                INNER JOIN [AR].[File] F on F.FileID = PF.FileID
                INNER JOIN [AR].[FileType] FT on FT.FileTypeID = F.FileTypeID
                WHERE P.EnvelopeID = {id}";
            }
            else
            {   
                if (!enveloperange.Contains('-'))
                {
                    Console.WriteLine("Envelope range must be in the format start-end.  Use the -h option for help.");
                    return;
                }
                var tokens = enveloperange.Split('-');
                if (tokens.Length != 2)
                {
                    Console.WriteLine("Envelope range must be in the format start-end.  Use the -h option for help.");
                    return;
                }
                if (!int.TryParse(tokens[0], out int start))
                {
                    Console.WriteLine("Envelope range must be in the format start-end and start must be an integer.  Use the -h option for help.");
                    return;
                }
                if (!int.TryParse(tokens[1], out int end))
                {
                    Console.WriteLine("Envelope range must be in the format start-end and end must be an integer.");
                    return;
                }

                query = $@"
                SELECT P.EnvelopeID, P.Sequence, FT.Description as Extension, F. Content 
                FROM [AR].[Page] P 
                INNER JOIN [AR].[PageFile] PF on PF.PageID = P.PageID 
                INNER JOIN [AR].[File] F on F.FileID = PF.FileID
                INNER JOIN [AR].[FileType] FT on FT.FileTypeID = F.FileTypeID
                WHERE P.EnvelopeID between {start} and {end}";
            }

            string connectionString = $"Server={server};Database={db};Trusted_Connection=True;";

            using (SqlConnection connection = new SqlConnection(connectionString))
            {
                SqlCommand command = new SqlCommand(query, connection);

                connection.Open();

                SqlDataReader reader = command.ExecuteReader();

                while (reader.Read())
                {
                    var envelopeId = (long)reader[0];
                    var sequence = (short)reader[1];
                    var extension = (string)reader[2];
                    var data = (byte[])reader[3];
                    File.WriteAllBytes($"{envelopeId}_{sequence}.{extension}", data);
                }
            }
        });

        return await rootCommand.InvokeAsync(args);
    }
}